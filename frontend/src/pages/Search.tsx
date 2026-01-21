import { Link } from "react-router-dom";
import { useSearchParams } from "react-router-dom";

export default function Search() {
    const [searchParams] = useSearchParams();
    const query = searchParams.get("query") ?? "";

    return (
        <div className="min-h-screen bg-(--color-bg) text-(--color-text) grid grid-rows-[auto_1fr_auto]">
            <header className="w-full flex justify-center border-t-8 border-t-(--color-primary) border-b-2 border-b-black/5 bg-white">
                <div className="w-full max-w-6xl flex items-center justify-between px-6 py-5">
                    <h1 className="text-2xl font-bold tracking-tight text-(--color-text-muted)"><Link to="/">OnTapAppRG</Link></h1>
                </div>
            </header>
            <main className="flex-1 w-full flex justify-center">
                <div className="w-full max-w-6xl px-6 py-8">
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
                    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                        {/* BeerCard entra aqui */}
                        <p className="text-(--color-text-muted)">Showing results for "<strong>{query}</strong>"</p>
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