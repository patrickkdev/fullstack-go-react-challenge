import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'

import { createBrowserRouter } from "react-router";
import { RouterProvider } from "react-router/dom";

import LandingPage from "./pages/landing-page"

const router = createBrowserRouter([
  {
    path: "/",
    element: LandingPage,
      { index: true, Component: Home },
  { path: "about", Component: About },
  {
    path: "auth",
    Component: AuthLayout,
    children: [
      { path: "login", Component: Login },
      { path: "register", Component: Register },
    ],
  },
  },
]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
