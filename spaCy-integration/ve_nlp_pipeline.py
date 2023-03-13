import json
import spacy
import random
from spacy.tokens import Span, DocBin
import os


def getWordIndex(text, charIndex):
    wordArray = text.split()
    count = 0
    for index, word in enumerate(wordArray):
        count += len(word)
        if count > charIndex:
            return index
        count += 1


nlp = spacy.blank("en")

with open('./output/json/testFile.jsonl', 'r') as json_file:
    json_list = list(json_file)
    # doc_list=[None for _ in range(len(json_list))]
    doc_list = []

for json_str in json_list:
    i = 0
    try:
        result = json.loads(json_str)
    except:
        continue
    # print(f"result: {result}")
    try:
        doc_list.append(nlp(result['text']))
        doc_list[i].ents = [Span(doc_list[i], getWordIndex(result['text'], result['spans'][0]['start']),
                                 getWordIndex(result['text'], result['spans'][0]['start']), label=result['spans'][0]['label'])]
        # print(f"{result['text']} {getWordIndex(result['text'], result['spans'][0]['start'])} {getWordIndex(result['text'], result['spans'][0]['start'])}")
    except IndexError as e:
        # print(result['text'])
        # print(e)
        continue
    i = i+1
    # print(f"result: {result}")
    # print(isinstance(result, dict))
    # getWordIndex(result['text'], result['spans'][0]['start'])
# print(f"doc_list: {doc_list}")
random.shuffle(doc_list)
train_docs = doc_list[:(len(doc_list)//2)]
dev_docs = doc_list[(len(doc_list)//2):]
# Create and save a collection of training docs
train_docbin = DocBin(docs=train_docs)
train_docbin.to_disk("./train.spacy")
# Create and save a collection of evaluation docs
dev_docbin = DocBin(docs=dev_docs)
dev_docbin.to_disk("./dev.spacy")

os.chdir("./spaCy-integration")

os.system("python3 -m spacy train ./config.cfg --output ./output --paths.train train.spacy --paths.dev dev.spacy ")

# $ python3 -m spacy package ./output/model-best ./packages --name ve_nlp_pipeline --version 1.0.0
# $ cd ./packages/en_my_pipeline-1.0.0
# $ pip install dist/en_my_pipeline-1.0.0.tar.gz

os.system("python3 -m spacy package ./output/model-best ./packages --name ve_nlp_pipeline --version 1.0.0 --force")
os.system(
    "pip install ./packages/en_ve_nlp_pipeline-1.0.0/dist/en_ve_nlp_pipeline-1.0.0.tar.gz")
