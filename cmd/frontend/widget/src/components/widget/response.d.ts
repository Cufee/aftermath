interface WidgetData {
  account: Account;
  cards: Cards;
  options: Options;
}

interface Options {
  flavor: string;
  rating: CardStyle;
  unrated: CardStyle;
  vehicles: VehiclesStyle;
}

interface CardStyle {
  visible: boolean;
  showTitle: boolean;
  showCareer: boolean;
  showLabel: boolean;
  blocks: string[] | null;
}
interface VehiclesStyle extends CardStyle {
  limit: number;
}

interface Cards {
  unrated: Unrated;
  rating: Rating;
}

interface Rating {
  overview: Overview;
}

interface Unrated {
  overview: Overview;
  vehicles: Vehicle[];
  highlights: Vehicle[];
}

interface Vehicle {
  type: string;
  title: string;
  blocks: Block[];
  meta: string;
}

interface Overview {
  type: string;
  title: string;
  blocks: BlockColumn[];
}

interface BlockColumn {
  blocks: Block[];
  flavor: string;
}

interface Block {
  data: Data;
  tag: string;
  label: string;
  value: string;
  meta?: string;
}

interface Data {
  session: string;
  career: string;
}

interface Account {
  id: string;
  realm: string;
  nickname: string;
  private: boolean;
  createdAt: string;
  lastBattleTime: string;
  clanId: string;
  clanTag: string;
}
