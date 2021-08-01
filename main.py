import json
import csv
import requests
import io


def fetch_csv_file(CSV_URL: str):
    r = requests.get(CSV_URL)
    r.encoding = 'cp932'
    csvio = io.StringIO(r.text, newline="")
    return csvio


def parse_CSV_to_JSON(csv_data: io.StringIO) -> dict:
    json_data = {}
    json_list = []

    for row in csv.DictReader(csv_data):
        json_list.append(row)
    json_data["data"] = json_list
    return json_data


def load_JSON(json_data: dict, json_file: str):
    with open(json_file, 'w', encoding="utf-8") as f:
        json.dump(json_data, f, ensure_ascii=False, indent=2)


def script_load_positive_cases():
    CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensya.csv"
    csv_data = fetch_csv_file(CSV_URL)
    json_data = parse_CSV_to_JSON(csv_data)
    json_file = "data/kansensya.json"
    load_JSON(json_data, json_file)
    return


def script_load_positivity_rate_in_testing():
    CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kensa.csv"
    csv_data = fetch_csv_file(CSV_URL)
    json_data = parse_CSV_to_JSON(csv_data)
    json_file = "data/kensa.json"
    load_JSON(json_data, json_file)
    return


def script_load_details_of_test_positives():
    CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensyazokusei.csv"
    csv_data = fetch_csv_file(CSV_URL)
    json_data = parse_CSV_to_JSON(csv_data)
    json_file = "data/kansensyazokusei.json"
    load_JSON(json_data, json_file)
    return


def main():
    script_load_details_of_test_positives()
    script_load_positive_cases()
    script_load_positivity_rate_in_testing()


if __name__ == "__main__":
    main()
