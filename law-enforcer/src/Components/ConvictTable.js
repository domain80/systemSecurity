// src/components/ConvictTable.js
import React from 'react';
import './ConvictTable.css';

function ConvictTable({ data }) {
  return (
    <table className="convict-table">
      <thead>
        <tr>
          <th>Convicts</th>
          <th>Arrested</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item, index) => (
          <tr key={index}>
            <td>{item.name}</td>
            <td>{item.arrested ? 'True' : 'False'}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default ConvictTable;
