export interface ResDepartment {
  department: Department[];
}

interface Department {
  id: number;
  name: string;
  number: string;
  information: Information;
  item: number[];
}

interface Information {
  area: string;
  text: string;
  access: string;
  address1: string;
  address2: string;
  address3: string;
  homepage: string;
  postcode: string;
  phone_number: string;
  displayed_name: string;
  nearest_station: string[];
  displayed_name_kana: string;
}
