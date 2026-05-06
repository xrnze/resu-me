import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import {
  createRouter,
  createRoute,
  createRootRoute,
  RouterProvider,
} from "@tanstack/react-router";
import "./index.css";
import { RootLayout } from "./routes/__root";
import { LandingPage } from "./routes/index";
import { AnalyzePage } from "./routes/analyze";

const rootRoute = createRootRoute({
  component: RootLayout,
});

const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: LandingPage,
});

const analyzeRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/analyze",
  component: AnalyzePage,
});

const routeTree = rootRoute.addChildren([indexRoute, analyzeRoute]);

const router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

function App() {
  return (
    <StrictMode>
      <RouterProvider router={router} />
    </StrictMode>
  );
}

createRoot(document.getElementById("root")!).render(<App />);
