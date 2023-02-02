import { createBrowserRouter, makeRouteConfig, Route } from "found";
import { Home } from "./home/Home";
import Layout from "./Layout";

export const BrowserRouter = createBrowserRouter({
  routeConfig: makeRouteConfig(
    <>
      <Route path="/" Component={Layout}>
        <Route Component={Home} />
      </Route>
    </>
  ),
});
