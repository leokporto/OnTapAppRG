import type { Beer } from '../types/Beer'

export async function fetchBeers(): Promise<Beer[]> {
  const response = await fetch('/api/beers')

  if (!response.ok) {
    throw new Error('Failed to fetch beers')
  }

  return response.json()
}

