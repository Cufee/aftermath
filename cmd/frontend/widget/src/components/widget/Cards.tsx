import { h } from '@stencil/core';

export default ({ data }: { data: WidgetData | null }) => {
  if (!data) {
    return <div class="bg-indigo-500 p-6 rounded-md flex justify-center">invalid cards data</div>;
  }
  return (
    <div class="flex flex-col gap-2">
      <OverviewCard data={data.cards.rating.overview} />
      <OverviewCard data={data.cards.unrated.overview} />
      {data.cards.unrated.vehicles.map(vehicle => (
        <VehicleCard data={vehicle} />
      ))}
    </div>
  );
};

const OverviewCard = ({ data }: { data: Overview }) => {
  return (
    <div class="flex flex-col gap-1 card overview-card grow">
      <span class="text-center text-gray-300 title">{data.title}</span>
      <div class="columns overview-columns flex flex-row gap-1 items-center bg-black rounded-xl bg-opacity-80 p-4">
        {data.blocks.map(column => (
          <Column column={column} />
        ))}
      </div>
    </div>
  );
};

const VehicleCard = ({ data }: { data: Vehicle }) => {
  return <div>{data.title}</div>;
};

const Column = ({ column }: { column: BlockColumn }) => {
  if (['rating', 'wn8'].includes(column.flavor)) {
    return 'special';
  }
  return (
    <div class="flex flex-col items-center justify-center column overview-column gap-2 grow">
      {column.blocks.map(block => (
        <Block block={block} />
      ))}
    </div>
  );
};

const Block = ({ block }: { block: Block }) => {
  return (
    <div class="flex flex-col items-center justify-between block">
      <span class="text-2xl text-white">{block.data.session}</span>
      {block.data.career && <span class="text-xl text-gray-300">{block.data.career}</span>}
      {block.label && <span class="text-sm text-gray-600">{block.label}</span>}
    </div>
  );
};
