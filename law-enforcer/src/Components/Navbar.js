// src/components/Navbar.js
import React from 'react';
import './Navbar.css';

function Navbar() {
  return (
    <nav className="navbar">
      <div className="navbar-brand">Law Enforcer</div>
      <div className="user-profile">
        <img src="https://via.placeholder.com/30" alt="User" className="user-avatar" />
        <span className="user-name">Eugene Cobbah</span>
      </div>
      
    </nav>
  );
}

export default Navbar;
