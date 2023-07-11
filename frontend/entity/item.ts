import { Kind } from './kind';

export type Item = {
  id: number;
  name: string;
  price: number;
  remarks: string;
  kinds: Kind[];
}
