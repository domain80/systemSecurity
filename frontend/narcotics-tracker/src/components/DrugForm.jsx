import React, { useState } from 'react';

const DrugForm = () => {
  const [formData, setFormData] = useState({
    convictName: '',
    serialNo: '',
    status: 'confiscated',
    tagId: 'NCB-DRUG-2024-001',
    verdict: 'approved',
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Form data submitted:', formData);
  };

  return (
    <form onSubmit={handleSubmit} className='log-form border w-[500px] h-[500px] content-evenly flex flex-col items-center justify-center p-5 rounded bg-slate-200'>
        <h1 className='font-bold font-sans text-xl'>Enter Details of Confiscated Drugs</h1>
      <div className='my-5'>
        <label>
          Convict Name:
          <input
            type="text"
            name="convictName"
            value={formData.convictName}
            onChange={handleChange}
             className="border"
            required
          />
        </label>
      </div>
      <div  className='my-5'>
        <label>
          Serial No:
          <input
            type="text"
            name="serialNo"
            className="border"
            value={formData.serialNo}
            onChange={handleChange}
            required
          />
        </label>
      </div>
      <div className='my-5'>
        <label>
          Status:
          <select name="status" value={formData.status} onChange={handleChange}>
            <option value="confiscated">Confiscated</option>
            <option value="destroyed">Destroyed</option>
          </select>
        </label>
      </div>
      <div className='my-5'>
        <label>
          Tag ID:
          <select name="tagId" value={formData.tagId} onChange={handleChange}>
            <option value="NCB-DRUG-2024-001">NCB-Hallucinogens</option>
            <option value="NCB-DRUG-2024-002">NCB-Inhalants</option>
            <option value="NCB-DRUG-2024-003">NCB-Stimulants</option>
          </select>
        </label>
      </div>
      <div className='my-5'>
        <label>
          Verdict:
          <select name="verdict" value={formData.verdict} onChange={handleChange}>
            <option value="approved">Approved</option>
            <option value="rejected">Rejected</option>
            <option value="pending">Pending</option>
          </select>
        </label>
      </div>
      <button type="submit" className='bg-blue-600 w-44 h-10 rounded text-white'>Submit</button>
    </form>
  );
};

export default DrugForm;
