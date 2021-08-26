import main
# import sys
import os


CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensya.csv"
file_name = "data/kansensya.json"
main.load_script(csv_url=CSV_URL, file_name=file_name)

CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kensa.csv"
file_name = "data/kensa.json"
main.load_script(csv_url=CSV_URL, file_name=file_name)

CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensyazokusei.csv"
file_name = "data/kansensyazokusei.json"
main.load_script(csv_url=CSV_URL, file_name=file_name)

CSV_URL = "https://www.city.kobe.lg.jp/documents/36421/kansensyazokusei.csv"
file_name = "data/36421-kansensyazokusei.json"
main.load_script(csv_url=CSV_URL, file_name=file_name)

# args = sys.argv
# base_url = args[1]

base_url = "https://api-cache.vaccines.sciseed.jp/public/281000/"

urls = [
    base_url + "articles/",
    base_url + "department/",
    base_url + "available_department/",
    base_url + "item/",
    base_url + "available_date/?department_id=8769&item_id=3&year=2021&month=8",
    base_url + "available_date/?department_id=8769&item_id=3&year=2021&month=9",
    base_url + "reservation_frame/?department_id=8769&item_id=3&start_date_after=2021-07-14&start_date_before=2022-10-10",
    base_url + "reservation_frame/?department_id=8770&item_id=3&start_date_after=2021-07-14&start_date_before=2022-10-10"
]
file_names = [
    "data/articles.json",
    "data/department.json",
    "data/available_department.json",
    "data/item.json",
    "data/month-8-all.json",
    "data/month-9-all.json",
    "data/reservation_frame1.json",
    "data/reservation_frame2.json"
]

url = base_url + "articles/"
file_name = "data/articles.json"
main.fetch_load_json(URL=url, file_path=file_name)

url = base_url + "department/"
file_name = "data/department.json"
main.fetch_load_json(URL=url, file_path=file_name)

url = base_url + "available_department/"
file_name = "data/available_department.json"
main.fetch_load_json(URL=url, file_path=file_name)

url = base_url + "item/"
file_name = "data/item.json"
main.fetch_load_json(URL=url, file_path=file_name)

url = base_url + "available_date/?department_id=8769&item_id=3&year=2021&month=8"
file_name = "data/month-8-all.json"
main.fetch_load_json(URL=url, file_path=file_name)

url = base_url + "available_date/?department_id=8769&item_id=3&year=2021&month=9"
file_name = "data/month-9-all.json"
main.fetch_load_json(URL=url, file_path=file_name)

url = base_url + "reservation_frame/?department_id=8769&item_id=3&start_date_after=2021-07-14&start_date_before=2022-10-10"
file_name = "data/reservation_frame1.json"
main.fetch_load_json(URL=url, file_path=file_name)

url = base_url + "reservation_frame/?department_id=8770&item_id=3&start_date_after=2021-07-14&start_date_before=2022-10-10"
file_name = "data/reservation_frame2.json"
main.fetch_load_json(URL=url, file_path=file_name)






# for i in len(urls):
#     main.fetch_load_json(URL=urls[i], file_path=file_names[i])
