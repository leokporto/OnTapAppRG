import React from "react";
import { useSearchParams } from "react-router-dom";
import type { Beer } from "../types/Beer";
import BeerCard from "../components/BeerCard";


export default function Search() {
    const [searchParams] = useSearchParams();
    const initialQuery = searchParams.get("query") ?? "";
    const [query, setQuery] = React.useState<string>(initialQuery);
    const [lastQuery, setLastQuery] = React.useState<string>(initialQuery);
    const [beers, setBeers] = React.useState<Beer[]>([]);
    const [loading, setLoading] = React.useState<boolean>(false);
    const [error, setError] = React.useState<string | null>(null);
    const [beerCount, setBeerCount] = React.useState<number>(0);

    async function fetchBeers(term: string): Promise<void> {
        if (!term.trim()) {
            setBeers([]);
            setBeerCount(0);
            return;
        }

        setLoading(true);
        setError(null);

        try {
            const response = await fetch(
            `http://localhost:8080/api/beers?fname=${encodeURIComponent(term)}`
            );

            if (!response.ok) {
                throw new Error(`HTTP ${response.status}`);
            }

            const data: Beer[] = await response.json();

            if(!(data===null || data.length===0)){
                setBeers(data);
                setBeerCount(data.length);
            } 
            else {
                setBeers([]);
                setBeerCount(0);
            }
        } catch (err) {
            setError("Failed to load beers. " + err);
            setBeers([]);
            setBeerCount(0);
        } finally {
            setLoading(false);
        }
    }

    React.useEffect(() => {
        if (initialQuery.trim() === "") {
            setBeers([]);
            return;
        }

        fetchBeers(initialQuery);
        setLastQuery(initialQuery);
    }, [initialQuery]);

    function handleSearchClick(): void {
        fetchBeers(query);
        setLastQuery(query);
    }

    return (
        <>
            <div className="flex gap-6 mb-10">
                <input type="text" value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    className="flex-1 px-6 py-4 text-lg                            
                    bg-(--color-surface)
                    text-(--color-text-muted)
                    placeholder:text-(--color-text-muted)
                    shadow-sm
                    focus:outline-none
                    focus:ring-2
                    focus:ring-(--color-primary)" />
                <button className="px-6 py-4 text-lg font-medium
                    rounded-sm
                    bg-(--color-accent)
                    text-(--color-text-muted)
                    hover:bg-(--color-accent)/80
                    hover:cursor-pointer
                    transition-colors
                    shadow-md
                    disabled:opacity-50
                    disabled:cursor-not-allowed"
                    onClick={handleSearchClick}>
                        Search
                </button>
            </div>        

            {/* Results placeholder */}
            <div className="grid grid-cols-1">
                {/* BeerCard entra aqui */}
                <p className="text-(--color-text-muted)">Showing {beerCount} results for "<strong>{lastQuery}</strong>"</p>
                <br/>
                {loading && <p className="text-(--color-text-muted)">Loading...</p>}
                {error && <p className="text-(--color-danger)">{error}</p>}                        
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