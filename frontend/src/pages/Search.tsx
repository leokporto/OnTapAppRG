import React from "react";
import { Link } from "react-router-dom";
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
        <div className="min-h-screen bg-(--color-bg) text-(--color-text) grid grid-rows-[auto_1fr_auto]">
            <header className="w-full flex justify-center border-t-8 border-t-(--color-primary) border-b-2 border-b-black/5 bg-white">
                <div className="w-full max-w-6xl flex items-center justify-between px-6 py-5">
                    <h1 className="text-2xl flex-col font-bold tracking-tight text-(--color-text-muted)"><Link to="/">OnTapAppRG</Link></h1>

                    <div className="flex-col w-lg" ></div>
                    <div className="flex-col-reverse w-fit">
                        <ul className="flex items-end  gap-8">
                            
                            <li className="mr-6 text-2md font-bold text-(--color-text-muted)"><Link to="/about">About</Link></li>
                            <li className="text-2md font-bold text-(--color-text-muted)"><Link to="/login">Login</Link></li>
                            <li className="mr-6"></li>
                        </ul>
                    </div>
                    <button className="flex-col px-4 py-2 text-sm font-medium border border-(--color-primary)
                        rounded-full text-(--color-text-muted) hover:bg-(--color-primary) hover:cursor-pointer">Join now</button>
                      
                </div>
            </header>
            <main className="flex-1 w-full flex justify-center">
                <div className="w-full max-w-6xl px-6 py-8">
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
                    {/* Tabs */}
                    <div className="flex gap-6 border-b border-black/10 mb-6">
                        <button className="pb-2 border-b-2 border-(--color-primary) font-semibold">
                        Beers
                        </button>
                        <button className="pb-2 text-(--color-text-muted)">
                        Breweries
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
                </div>
            </main>
            <footer className="w-full flex justify-center py-6 bg-white">
                <div className="w-full max-w-6xl text-center px-6">
                    <p className="text-sm text-(--color-text-muted)">
                        Modern craft beer catalog 🍺
                    </p>
                </div>
            </footer>
        </div>
    );
}