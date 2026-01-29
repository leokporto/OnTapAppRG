import { useParams } from "react-router-dom";
import type { Beer } from "../types/Beer";
import React from "react";
import type { Brewery } from "../types/Brewery";
import BeerCard from "../components/BeerCard";
import brewery_icon from '../assets/brewery.png';


export default function BreweryDetails() {
    const { id } = useParams<{ id: string }>();
    const [beers, setBeers] = React.useState<Beer[]>([]);   
    const [brewery, setBrewery] = React.useState<Brewery | null>(null);
    const [error, setError] = React.useState<string | null>(null);

    React.useEffect(() => {
        if (!id) {
            setError("No brewery ID provided in URL.");
            return;
        }

        fetchBrewery(id);
    }, [id]);

    async function fetchBrewery(breweryId: string): Promise<void> {
        if (!breweryId) return;
        
        try{
            setError(null);

            const [beersResponse, breweryResponse] =  await Promise.all([
                fetch(`http://localhost:8080/api/beers?bid=${encodeURIComponent(breweryId)}`),
                fetch(`http://localhost:8080/api/breweries/${encodeURIComponent(breweryId)}`)
            ]);

            if(!breweryResponse.ok){
                throw new Error(`HTTP ${breweryResponse.status} when fetching brewery`);
            }

            if(!beersResponse.ok){
                throw new Error(`HTTP ${beersResponse.status} when fetching beers`);
            }

            const [beers, brewery]: [Beer[], Brewery] = await Promise.all([
                beersResponse.json(),
                breweryResponse.json()
            ]);
        
            setBeers(beers);
            setBrewery(brewery);
        } catch (err) {
            setError("Failed to load brewery details: " + err);
            setBeers([]);
            setBrewery(null);
        }
    }

    return (
    
        
        <>

            {error && <p className="text-red-500">{error}</p>}

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
                    <img src={brewery_icon} alt={`${brewery?.name} logo`} className='h-16 w-16 mx-auto' />
                </div>
                <div className="flex flex-col min-w-20">
                    <h3 className="text-lg font-bold text-(--color-text-muted) truncate">
                        {brewery?.name}
                    </h3>

                </div>
            </div>

            <br />

            <div className="grid grid-cols-1">
                <p className="text-(--color-text-muted)">Showing {beers.length} results for this brewery:</p>

                {beers.length === 0 && <p>No beers found for this brewery.</p>}
                {beers.length > 0 && (                       
                    <div className="grid grid-cols-1">
                        {beers.map((beer) => (
                        <BeerCard key={beer.id} beer={beer} />
                        ))}
                    </div>
                                            
                )}
            </div>
        </>
            
    );
}