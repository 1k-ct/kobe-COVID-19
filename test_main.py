import main

CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensya.csv"
file_name = "data/kansensya.json"
main.load_script(csv_url=CSV_URL, file_name=file_name)

CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensyazokusei.csv"
file_name = "data/kensa.json"
main.load_script(csv_url=CSV_URL, file_name=file_name)

CSV_URL = "https://www.city.kobe.lg.jp/documents/32576/kansensyazokusei.csv"
file_name = "data/kansensyazokusei.json"
main.load_script(csv_url=CSV_URL, file_name=file_name)
