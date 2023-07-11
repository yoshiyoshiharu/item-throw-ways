import { Kind } from './kind';

export type Item = {
  id: number;
  name: string;
  name_kana: string;
  price: number;
  remarks: string;
  kinds: Kind[];
}
