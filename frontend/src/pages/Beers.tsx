import { useEffect, useState } from "react"
import BeerList from "../components/BeerList"
import type { Beer } from "../types/Beer"
import { fetchBeers } from "../services/BeerSvc"

export default function Beers() {
    const [beers, setBeers] = useState<Beer[]>([]);
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        async function loadBeers() {
            try {
                setLoading(true)
                const data =  await fetchBeers()
                setBeers(data)
            } catch (error) {
                console.error("Error fetching beers:", error)
            } finally {
                setLoading(false)
            }
        }
        loadBeers()
    }, [])

    return (
        <div>
        <h2>Beers Page</h2>        
        {loading ? <p>Loading...</p> : <BeerList beers={beers} />}
        </div>
    )
}