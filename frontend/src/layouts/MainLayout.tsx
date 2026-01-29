import { Outlet, Link } from "react-router-dom";

export default function MainLayout() {
    return (
        <div className="min-h-screen bg-(--color-bg) text-(--color-text) grid grid-rows-[auto_1fr_auto]">
            <header className="w-full flex justify-center border-t-8 border-t-(--color-primary) border-b-2 border-b-black/5 bg-white">
                <div className="w-full max-w-6xl flex items-center justify-between px-6 py-5">
                    <h1 className="text-2xl flex-col font-bold tracking-tight text-(--color-text-muted)"><Link to="/">OnTapAppRG</Link></h1>

                    <div className="flex-col w-lg" >
                        
                    </div>                    
                    <div className="flex-col-reverse w-fit">
                        <ul className="flex items-end  gap-8">
                            <li className="mr-6 text-2md font-bold text-(--color-text-muted)"><Link to="/search">Search</Link></li>
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
                    <Outlet />
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