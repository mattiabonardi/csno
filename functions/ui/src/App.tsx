import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import GamePage from "./pages/GamePage";
import Home from "./pages/Home";
import NoPage from "./pages/NoPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
          <Route path="/" index element={<Home />} />
          <Route path="/game" element={<GamePage />} />
          <Route path="*" element={<NoPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
