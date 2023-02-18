import { createBrowserRouter, makeRouteConfig, Route } from "found";
import LoginPage from "../features/auth/login/LoginPage";
import { Home } from "./home/Home";
import Layout from "./Layout";

export const BrowserRouter = createBrowserRouter({
  routeConfig: makeRouteConfig(
    <>
      <Route path="/" Component={Layout}>
        <Route path="/login" Component={LoginPage} />
        <Route Component={Home} />
      </Route>
    </>
  ),
});
