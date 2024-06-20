// src/components/MainPage.js
import React from 'react';
import ConvictTable from './ConvictTable';
import './MainPage.css';

function MainPage() {
  // Sample data for the table
  const data = [
    { name: 'Bright Kumi', arrested: true },
    { name: 'Erica Brown', arrested: false },
    { name: 'Michael Skolo', arrested: true },
    { name: 'Shatawale', arrested: false },
    { name: 'Shatawale', arrested: false },
    { name: 'David Akama', arrested: true },
    { name: 'Divine Anum', arrested: true },
    { name: 'Cupid Mansayer', arrested: false },
    { name: 'Stoneboy', arrested: true },
    { name: 'Squad Brighton', arrested: false },
  ];

  return (
    <main className="main-page">
      <h2>List of People found with Drugs</h2>
      <ConvictTable data={data} />
    </main>
  );
}

export default MainPage;
