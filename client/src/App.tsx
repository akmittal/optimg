import React from "react";
import Dashboard from "./components/Dashboard";
import "./App.css";
import "@rmwc/icon/styles";
import { QueryCache, ReactQueryCacheProvider } from "react-query";

const queryCache = new QueryCache();



function App() {
  return (
    <ReactQueryCacheProvider queryCache={queryCache}>
      <div className="App">
        <Dashboard />
      </div>
    </ReactQueryCacheProvider>
  );
}

export default App;
