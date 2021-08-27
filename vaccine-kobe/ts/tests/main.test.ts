import { FetchItem } from "../src/app";
import fs from "fs";

const test = async () => {
  const url = "https://api-cache.vaccines.sciseed.jp/public/281000/";
  url + "department/";
  const vaccine = FetchItem<any>(url);
  // const vaccine = FetchItem<any>(url);
  try {
    const v = await vaccine;
    const jsonData = JSON.stringify(v, null, 4);
    fs.writeFileSync("../../data/department.json", jsonData);
  } catch (err) {
    console.log(err);
  }
};
test();
