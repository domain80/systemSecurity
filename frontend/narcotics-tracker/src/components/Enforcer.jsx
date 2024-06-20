import React from "react";
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
  { id: 5, lastName: 'Targaryen', firstName: 'Daenerys', Arrested: 0 },
  { id: 6, lastName: 'Melisandre', firstName: null, Arrested: 1},
  { id: 7, lastName: 'Clifford', firstName: 'Ferrara', Arrested: 1 },
  { id: 8, lastName: 'Frances', firstName: 'Rossini', Arrested: 1 },
  { id: 9, lastName: 'Roxie', firstName: 'Harvey', Arrested: 1 },
];
  return (
    <div className="h-max">

      <div className="body">
      <div style={{ height: 500, width: '100%' }}>
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


