import json
import spacy
import random
from spacy.tokens import Span, DocBin


def getWordIndex(text, charIndex):
    wordArray = text.split()
    count = 0
    for index, word in enumerate(wordArray):
        count += len(word)
        if count > charIndex:
            return index
        count += 1


nlp = spacy.blank("en")

with open('../output/json/testFile.jsonl', 'r') as json_file:
    json_list = list(json_file)
    # doc_list=[None for _ in range(len(json_list))]
    doc_list = []

for json_str in json_list:
    i = 0
    try:
        result = json.loads(json_str)
    except:
        continue
    print(f"result: {result}")
    doc_list.append(nlp(result['text']))
    doc_list[i].ents = [Span(doc_list[i], getWordIndex(result['text'], result['spans'][0]['start']),
                             getWordIndex(result['text'], result['spans'][0]['start']), label=result['spans'][0]['label'])]
    i = i+1
    # print(f"result: {result}")
    # print(isinstance(result, dict))
    # getWordIndex(result['text'], result['spans'][0]['start'])

print(f"doc_list: {doc_list}")
