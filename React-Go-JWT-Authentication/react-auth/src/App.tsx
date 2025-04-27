import React from 'react';
import './App.css';
import Login from "./pages/Login.tsx";
import Nav from "./components/Nav.tsx";
import Home from "./pages/Home.tsx";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Register from "./pages/Register.tsx";
import VerifyEmail from "./pages/verifyEmail.tsx";


function App() {
  return (
    <div className="App" style={{ backgroundColor: "grey" }}>
      <Nav />
      <main className="form-signin w-100 m-auto">
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/verify" element={<VerifyEmail />} />
          </Routes>
        </BrowserRouter>
      </main>
    </div>
  );
}

export default App;
