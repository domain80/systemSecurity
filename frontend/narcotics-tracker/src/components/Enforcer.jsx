import React from "react";

import { CgProfile } from "react-icons/cg";
import { DataGrid } from '@mui/x-data-grid';

const Enforcer = () => {
  

const columns = [
  { field: 'id', headerName: 'ID', width: 70 },
  { field: 'firstName', headerName: 'First name', width: 130 },
  { field: 'lastName', headerName: 'Last name', width: 130 },
  {
    field: 'Arrested',
    headerName: 'Arrested',
    type: 'boolean',
    width: 90,
  },
  {
    field: 'fullName',
    headerName: 'Full name',
    description: 'This column has a value getter and is not sortable.',
    sortable: false,
    width: 160,
    valueGetter: (value, row) => `${row.firstName || ''} ${row.lastName || ''}`,
  },
];

const rows = [
  { id: 1, lastName: 'Snow', firstName: 'Jon', Arrested: 1 },
  { id: 2, lastName: 'Lannister', firstName: 'Cersei', Arrested: 1 },
  { id: 3, lastName: 'Lannister', firstName: 'Jaime', Arrested: 1 },
  { id: 4, lastName: 'Stark', firstName: 'Arya', Arrested: 1 },
  { id: 5, lastName: 'Targaryen', firstName: 'Daenerys', Arrested: null },
  { id: 6, lastName: 'Melisandre', firstName: null, Arrested: 11 },
  { id: 7, lastName: 'Clifford', firstName: 'Ferrara', Arrested: 1 },
  { id: 8, lastName: 'Frances', firstName: 'Rossini', Arrested: 1 },
  { id: 9, lastName: 'Roxie', firstName: 'Harvey', Arrested: 1 },
];
  return (
    <div className="h-max">
      <div className="h-[60px] bg-slate-500 flex justify-between items-center px-11">
        <h1 className="text-white font-extrabold text-4xl">NCB</h1>
        <input
          className="h-[30px] w-7/12 rounded"
          type="search"
          placeholder="Search"
        />
        <div className="logoutnprofile flex">
          <button className="mx-3 text-white rounded bg-sky-700 w-[80px] h-[40px] mr-4">
            Logout
          </button>
          <CgProfile className="text-white my-3" />
        </div>
      </div>

      <div className="body ">
      <div style={{ height: 400, width: '100%' }}>
      <DataGrid
        rows={rows}
        columns={columns}
        initialState={{
          pagination: {
            paginationModel: { page: 0, pageSize: 5 },
          },
        }}
        pageSizeOptions={[5, 10]}
        checkboxSelection
      />
    </div>
      </div>
    </div>
  );
};

export default Enforcer;


