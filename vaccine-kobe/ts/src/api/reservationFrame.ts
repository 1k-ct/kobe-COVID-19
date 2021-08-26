export interface ReservationFrame {
  reservation_frame: Reservationframe[];
}

interface Reservationframe {
  id: number;
  name: string;
  start_at: string;
  end_at: string;
  is_published: boolean;
  reservation_cnt: number;
  department: number;
  reservation_cnt_limit: number;
  item: number;
  next?: any;
}
