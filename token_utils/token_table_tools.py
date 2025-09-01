import csv
import json

def csv_to_json(csv_path, json_path):
    with open(csv_path, 'r', encoding='utf-8') as csv_file:
        csv_reader = csv.DictReader(csv_file)

        data = []
        for row in csv_reader:
            converted = {}
            for key, value in row.items():
                if key in ['chain_id', 'token_id', 'decimals']:
                    converted[key] = int(value) if value.strip() else None
                else:
                    converted[key] = value
            data.append(converted)

    with open(json_path, 'w', encoding='utf-8') as json_file:
        json.dump(data, json_file, indent=2, ensure_ascii=False)

if __name__ == "__main__":
    csv_to_json('token_table.csv', 'output.json')
    print("outputfile: output.json")