import { RouterProvider, createBrowserRouter } from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import Search from "./pages/Search";

const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />,
  },
  {
    path: "/search",
    element: <Search />
  }
]);

export default function App() {
  return <RouterProvider router={router} />;
}
