import type { Beer } from '../types/Beer';

type BeerListProps = {
    beers: Beer[];
}
export default function BeerList({beers}: BeerListProps) {
    return (
        <table >
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Style</th>
                    <th>Brewery</th>
                    <th>ABV</th>
                    <th>IBU</th>
                </tr>
            </thead>
            <tbody>
                {beers.map((beer) => (
                    <tr key={beer.id}>
                        <td>{beer.name}</td>
                        <td>{beer.style}</td>
                        <td>{beer.brewery}</td>
                        <td>{beer.abv}%</td>
                        <td>{beer.minibu} - {beer.maxibu}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
}