import { RouterProvider, createBrowserRouter } from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import Search from "./pages/Search";
import BreweryDetails from "./pages/BreweryDetails";
import MainLayout from "./layouts/MainLayout";
import About from "./pages/About";

const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />,
  },
  { 
    element: <MainLayout />, 
    children: [
      {
        path: "/search",
        element: <Search />
      },
      {
        path: "/breweries/:id",
        element: <BreweryDetails />
      },
      {
        path: "/about",
        element: <About />
      }]
  }
]);

export default function App() {
  return <RouterProvider router={router} />;
}
