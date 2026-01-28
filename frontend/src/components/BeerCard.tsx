import type { Beer } from "../types/Beer"
import pint_icon from '../assets/pint.png';

type beerCardProps = {
    beer: Beer
}

export default function BeerCard({ beer }: beerCardProps) {
    return (
        <div className="
        bg-(--color-surface)
        rounded-lg
        shadow-sm
        p-4
        grid
        grid-cols-[64px_1fr_auto]
        gap-4
        items-center
        border border-black/5
        hover:shadow-md
        transition-shadow
      ">
        <div className="flex flex-col items-center justify-center min-w-4">
            <img src={pint_icon} alt="{beer.name} pint" className='rounded-full h-16 w-16 mx-auto' />
        </div>
        <div className="flex flex-col min-w-20">
            <h3 className="text-lg font-bold text-(--color-text-muted) truncate">
                {beer.name}
            </h3>

            <p className="text-sm text-(--color-text-muted) truncate">
                {beer.brewery}
            </p>

            <p className="text-sm italic text-(--color-text-muted) truncate">
                {beer.style}
            </p>
        </div>

        <div className="flex flex-col items-end justify-center gap-2">
            <span
            className="
                px-2 py-1
                text-xs
                font-semibold
                rounded
                bg-(--color-accent)
                whitespace-nowrap
                w-full
            ">
            ABV: {beer.abv.toFixed(1)}%
            </span>

            <span
            className="
                px-2 py-1
                text-xs
                font-semibold
                rounded
                bg-(--color-primary)
                whitespace-nowrap
                w-full"
            >
            IBU: {beer.minibu}–{beer.maxibu}
            </span>
        </div>
      </div>
    );
}