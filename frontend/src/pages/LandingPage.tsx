import React from "react";
import { useNavigate } from "react-router-dom";
import ontapapp_ico from '../assets/ontapapp.jpeg'

export default function LandingPage() {
    const [query, setQuery] = React.useState<string>("");
    const navigate = useNavigate();

    function handleSearch(evt: React.FormEvent<HTMLFormElement>): void {
        evt.preventDefault();
        
        const trimmed = query.trim();
        if(!trimmed) {
            return;
        }
        
        navigate(`/search?query=${encodeURIComponent(trimmed)}`);
    }

    return (
        <div className="min-h-screen bg-(--color-bg) text-(--color-text) grid grid-rows-[auto_1fr_auto]">            
            <header className="w-full flex justify-center bg-(--color-primary)">
                <div className="w-full max-w-6xl flex items-center justify-between px-6 py-5">
                    <h1 className="text-2xl font-bold tracking-tight text-(--color-text-muted)">OnTapAppRG</h1>

                    <button className="px-4 py-2 text-sm font-medium border border-(--color-text-muted)
                        rounded-md text-(--color-text-muted) hover:bg-(--color-surface) hover:cursor-pointer">Login</button>               
                </div>
            </header>
            <main className="flex-1 w-full flex justify-center">
                <div className="w-full max-w-6xl flex flex-col items-center justify-center px-6 text-center">
                    <h2 className="text-4xl sm:text-5xl font-extrabold mb-4 text-(--color-text-muted)">Discover amazing beers</h2>
                    
                    <img src={ontapapp_ico} alt="On Tap app logo" className='rounded-full h-64 w-64 mx-auto mb-6' />

                    <p className="text-(--color-text-muted) mb-10 max-w-xl">
                        Search beers, styles, and breweries in a modern and simple catalog.
                    </p>

                    <form onSubmit={handleSearch} className="w-full max-w-2xl">
                        <input type="text" value={query} onChange={(e) => setQuery(e.target.value)}
                            placeholder="Search for beers or breweries"
                            className="w-full px-6 py-5 text-lg
                            rounded-full
                            bg-(--color-surface)
                            text-(--color-text)
                            placeholder:text-(--color-text-muted)
                            shadow-sm
                            focus:outline-none
                            focus:ring-2
                            focus:ring-(--color-primary)" />
                    </form>
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