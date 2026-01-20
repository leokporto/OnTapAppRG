import './App.css'
import Beers from './pages/Beers'
import ontapapp_ico from './assets/ontapapp.jpeg'

function App() {

  return (
    <>
      <header>
        <img src={ontapapp_ico} alt="On Tap app logo" className='logo' /> 
        <h1>On Tap app (react + go)</h1>
      </header>
      <main>        
        <div>
          <Beers />
        </div>
      </main>
    </>

  )
}

export default App
