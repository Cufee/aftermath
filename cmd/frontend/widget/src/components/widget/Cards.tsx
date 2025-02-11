import { h } from '@stencil/core';

// Calculate the minimum width of columns in rem, balanced around the center column
const overviewColumnSizes = (overview: Overview & { style: CardStyle }, ...more: (Overview & { style: CardStyle })[]): Record<number, number> => {
  const columnSizes = {};
  const cards = [overview, ...more];
  for (const card of cards) {
    // cards have column that are aligned horizontally
    for (let i = 0; i < card.blocks.length; i++) {
      if (!columnSizes[i]) columnSizes[i] = 0;
      for (const block of card.blocks[i].blocks) {
        // column have blocks that are aligned vertically
        const contentLength = Math.max(
          // blocks have content that is aligned vertically
          card.style.showCareer ? block.data.career.length * 0.75 : 0, // career text is smaller
          card.style.showLabel ? block.label.length * 0.5 : 0, // label text is smaller
          block.data.session.length,
        );
        if (columnSizes[i] < contentLength) {
          columnSizes[i] = contentLength;
        }
      }
    }
  }

  const balancedSizes = {};
  const columnCount = Object.keys(columnSizes).length;
  for (let i = 0; i < Math.ceil(columnCount / 2); i++) {
    const mirroredIndex = columnCount - 1 - i;
    const maxSize = Math.max(columnSizes[i] || 0, columnSizes[mirroredIndex] || 0);
    balancedSizes[i] = maxSize;
    if (i !== mirroredIndex) {
      balancedSizes[mirroredIndex] = maxSize;
    }
  }

  return Object.assign(columnSizes, balancedSizes);
};

const vehicleBlockSizes = (vehicles: Vehicle[], style: VehiclesStyle): Record<number, number> => {
  const columnSizes = {};
  for (const vehicle of vehicles) {
    // vehicles have blocks that are aligned horizontally
    for (let i = 0; i < vehicle.blocks.length; i++) {
      if (!columnSizes[i]) columnSizes[i] = 0;
      const contentLength = Math.max(
        // blocks have content that is aligned vertically
        style.showCareer ? vehicle.blocks[i].data.career.length * 0.75 : 0,
        style.showLabel ? vehicle.blocks[i].label.length * 0.5 : 0,
        vehicle.blocks[i].data.session.length,
      );
      if (columnSizes[i] < contentLength) {
        columnSizes[i] = contentLength;
      }
    }
  }
  return columnSizes;
};

export default ({ cards, options, assetsDomain }: { cards: Cards | null; options: Options | null; assetsDomain: string }) => {
  if (!cards) {
    return <div>invalid cards data</div>;
  }
  if (!options) {
    return <div>invalid options data</div>;
  }

  const overviewSizes = overviewColumnSizes({ ...cards.rating.overview, style: options.rating }, { ...cards.unrated.overview, style: options.unrated });
  const vehicleSizes = vehicleBlockSizes(cards.unrated.vehicles || [], options.vehicles);

  return (
    <div class="flex flex-col gap-2 min-w-fit">
      {options.unrated.visible && <OverviewCard data={cards.unrated.overview} style={options.unrated} columnSizes={overviewSizes} assetsDomain={assetsDomain} />}
      {options.rating.visible && <OverviewCard data={cards.rating.overview} style={options.rating} columnSizes={overviewSizes} assetsDomain={assetsDomain} />}
      {options.vehicles.visible &&
        cards.unrated.vehicles
          .slice(0, options.vehicles.limit)
          .map(vehicle => <VehicleCard data={vehicle} style={options.vehicles} blockSizes={vehicleSizes} assetsDomain={assetsDomain} />)}
    </div>
  );
};

const OverviewCard = ({ data, style, columnSizes, assetsDomain }: { data: Overview; style: CardStyle; columnSizes: Record<number, number>; assetsDomain: string }) => {
  return (
    <div class="flex flex-col gap-2 card overview-card grow">
      {style.showTitle && <span class="text-center text-gray-300 title">{data.title}</span>}
      <div class="columns overview-columns justify-around flex flex-row gap-2 items-center bg-black rounded-3xl bg-opacity-80 p-4">
        {data.blocks.map((column, i) => (
          <Column column={column} style={style} width={columnSizes[i]} assetsDomain={assetsDomain} />
        ))}
      </div>
    </div>
  );
};

const VehicleCard = ({ data, style, blockSizes, assetsDomain }: { data: Vehicle; style: VehiclesStyle; blockSizes: Record<number, number>; assetsDomain: string }) => {
  const css = (i: number) => {
    return { 'min-width': `${blockSizes[i] || 0}em` };
  };
  const iconPath = data.blocks.find(b => ['wn8', 'ranked_rating'].includes(b.tag)).meta;

  return (
    <div class="flex flex-col gap-2 card vehicle-card grow bg-black rounded-3xl bg-opacity-80 py-4 px-6">
      {style.showTitle && (
        <div class="flex flex-row gap-2 justify-between">
          <span class="text-gray-300 title">{data.title}</span>
          {iconPath && <img src={assetsDomain + iconPath} class="size-5" />}
        </div>
      )}
      <div class="blocks vehicle-blocks flex flex-row gap-2 items-center justify-around">
        {data.blocks.map((block, i) => (
          <div style={css(i)}>
            <Block block={block} style={style} />
          </div>
        ))}
      </div>
    </div>
  );
};

const Column = ({ column, style, width, assetsDomain }: { column: BlockColumn; style: CardStyle; width: number; assetsDomain: string }) => {
  const css: Record<string, string> = {};
  css['min-width'] = `${width || 0}em`;

  if (['rating', 'wn8'].includes(column.flavor)) {
    const iconPath = column.blocks.find(b => ['wn8', 'ranked_rating'].includes(b.tag)).meta;

    return (
      <div class="flex flex-col items-center justify-center column overview-column special-overview-column gap-1" style={css}>
        {iconPath && <img src={assetsDomain + iconPath} class="size-16" />}
        {column.blocks.map(block => (
          <Block block={block} style={style} />
        ))}
      </div>
    );
  }

  const blocks = style.blocks?.length > 0 ? column.blocks.filter(b => style.blocks.includes(b.tag)) : column.blocks;
  return (
    <div class="flex flex-col items-center justify-center column overview-column gap-2" style={css}>
      {blocks.map(block => (
        <Block block={block} style={style} />
      ))}
    </div>
  );
};

const Block = ({ block, style }: { block: Block; style: CardStyle }) => {
  return (
    <div class="flex flex-col items-center justify-between block text-nowrap">
      <span class="text-2xl text-white">{block.data.session}</span>
      {style.showCareer && <span class="text-xl text-gray-300">{block.data.career}</span>}
      {style.showLabel && <span class="text-sm text-gray-600">{block.label}</span>}
    </div>
  );
};
