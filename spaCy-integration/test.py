import json

with open('../output/json/testFile.jsonl', 'r') as json_file:
    json_list = list(json_file)

for json_str in json_list:
    try:
        result = json.loads(json_str)
    except  : 
        continue
    
        
    print(f"result: {result}")
    print(isinstance(result, dict))
