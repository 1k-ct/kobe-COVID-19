export interface ResItem {
  item: Item[];
}

interface Item {
  id: number;
  name: string;
  interval: number;
  information: Information;
}

interface Information {
  message: string;
  displayed_name_kana: string;
  vaccine_manufacturer: string;
}
