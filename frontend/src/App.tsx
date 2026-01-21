import ontapapp_ico from './assets/ontapapp.jpeg'

function App() {

  return (
    <>
      <header className='font-[--font-sans]'>
        <img src={ontapapp_ico} alt="On Tap app logo" className='rounded-full h-24 w-24' /> 
        <h1>On Tap app (react + go)</h1>
      </header>
      <main className='bg-[--color-surface] font-[--font-sans]'>        
        <div>
          App INiciado
        </div>
      </main>
    </>

  )
}

export default App
