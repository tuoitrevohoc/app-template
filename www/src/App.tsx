import { Suspense, useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";
import { RelayEnvironmentProvider } from "react-relay";
import environment from "./relay/environment";
import ErrorBoundary from "./features/error-handling/ErrorBoundary";
import { BrowserRouter } from "./routes";

function App() {
  const [count, setCount] = useState(0);

  return (
    <RelayEnvironmentProvider environment={environment}>
      <ErrorBoundary>
        <Suspense fallback={<div>Loading...</div>}>
          <BrowserRouter />
        </Suspense>
      </ErrorBoundary>
    </RelayEnvironmentProvider>
  );
}

export default App;
