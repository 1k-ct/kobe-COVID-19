import json
import csv
import requests
import io
import os
# from os.path import join, dirname
# from dotenv import load_dotenv


def parse_CSV_to_JSON(CSV_URL: str) -> dict:
    json_data = {}
    json_list = []

    r = requests.get(CSV_URL)
    r.encoding = 'cp932'
    csv_data = io.StringIO(r.text, newline="")

    for row in csv.DictReader(csv_data):
        json_list.append(row)
    json_data["data"] = json_list
    return json_data


def valid_kansesya_json(json_data: dict):
    # FIXME とりあえずのデータの行数が2000行以上なら有効
    # データの内容が有効か
    return len(json_data["data"]) < 600


def load_JSON(json_data: dict, json_file: str):
    if valid_kansesya_json(json_data):
        return
    with open(json_file, 'w', encoding="utf-8") as f:
        json.dump(json_data, f, ensure_ascii=False, indent=4)


def load_script(csv_url: str, file_name: str) -> None:
    json_data = parse_CSV_to_JSON(csv_url)
    load_JSON(json_data, file_name)
    return


def validate_kansensya_csv_file(csv_file):
    # 取得したCSVファイルを正確な内容か検証する
    pass


# def setting():
#     load_dotenv(verbose=True)

#     dotenv_path = join(dirname(__file__), ".env")
#     load_dotenv(dotenv_path)
#     return os.environ.get("URL")


def fetch_load_json(URL: str, file_path: str):
    r = requests.get(URL)

    data = r.json()
    with open(file_path, 'w', encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=4)


def main():

    # CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensya.csv"
    # file_name = "data/kansensya.json"
    # load_script(csv_url=CSV_URL, file_name=file_name)

    # CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kensa.csv"
    # file_name = "data/kensa.json"
    # load_script(csv_url=CSV_URL, file_name=file_name)

    # CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensyazokusei.csv"
    # file_name = "data/kansensyazokusei.json"
    # load_script(csv_url=CSV_URL, file_name=file_name)

    # CSV_URL = "https://www.city.kobe.lg.jp/documents/36421/kansensyazokusei.csv"
    # file_name = "data/36421-kansensyazokusei.json"
    # load_script(csv_url=CSV_URL, file_name=file_name)
    url = "https://api-cache.vaccines.sciseed.jp/public/281000/articles/"
    file_name = "data/article.json"
    fetch_load_json(URL=url, file_path=file_name)

    url = "https://api-cache.vaccines.sciseed.jp/public/281000/department/"
    file_name = "data/department.json"
    fetch_load_json(URL=url, file_path=file_name)

    url = "https://api-cache.vaccines.sciseed.jp/public/281000/available_department/"
    file_name = "data/available_department.json"
    fetch_load_json(URL=url, file_path=file_name)

    url = "https://api-cache.vaccines.sciseed.jp/public/281000/item/"
    file_name = "data/item.json"
    fetch_load_json(URL=url, file_path=file_name)

    url = "https://api-cache.vaccines.sciseed.jp/public/281000/available_date/?department_id=8769&item_id=3&year=2021&month=8"
    # BASE_URL = setting()
    # URL = f"{BASE_URL}/available_date/?department_id=8769&item_id=3&year=2021&month=8"
    # fetch_load_json(URL, "data/month-8-all.json")
    url = "https://api-cache.vaccines.sciseed.jp/public/281000/reservation_frame/?department_id=8769&item_id=3&start_date_after=2021-07-14&start_date_before=2022-10-10"
    file_name = "data/reservation_frame1.json"
    fetch_load_json(URL=url, file_path=file_name)
    url = "https://api-cache.vaccines.sciseed.jp/public/281000/reservation_frame/?department_id=8770&item_id=3&start_date_after=2021-07-14&start_date_before=2022-10-10"
    file_name = "data/reservation_frame2.json"
    fetch_load_json(URL=url, file_path=file_name)
    # BASE_URL = setting()
    # # start_date_after = "07-15"
    # start_date_after = "08-01"
    # start_date_before = "08-02"
    # URL = f"{BASE_URL}/reservation_frame/?department_id=8769&item_id=3&start_date_after=2021-{start_date_after}&start_date_before=2021-{start_date_before}"
    # fetch_load_json(URL, "data/date-details.json")
    print("ok")


if __name__ == "__main__":
    main()
