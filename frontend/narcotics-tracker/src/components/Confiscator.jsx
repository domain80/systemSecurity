import React from "react";
import { CgProfile } from "react-icons/cg";
import DrugForm from "./DrugForm";
import { useState } from "react";
import Enforcer from "./Enforcer";

const Confiscator = () => {
const [showDrugForm, setShowDrugForm] = useState(false);
const [allDrugs, setAllDrugs] = useState(false);

const handleShowDrugForm = () => {
    setShowDrugForm(true);
    setAllDrugs(false);
  };

  const handleShowAllDrugs = () => {
    setAllDrugs(true);
    setShowDrugForm(false);
  };

  return (
    <div className="h-max">
      <div className="top-nav h-[60px] bg-slate-500 flex justify-between items-center px-11 sticky top-0">
        <h1 className="text-white font-extrabold text-4xl">NCB</h1>
        <input
          className="h-[30px] w-7/12 rounded"
          type="search"
          placeholder="Search"
        />
        <div className="logoutnprofile flex">
          <button className="mx-3 text-white rounded-full bg-sky-700 w-[80px] h-[40px] mr-4">
            Logout
          </button>
          <CgProfile className="text-white my-3" />
        </div>
      </div>
      <div className="body flex items-center justify-between">
        <div className="side-nav border h-screen w-1/6 flex flex-col ">
          <button onClick={handleShowAllDrugs} className="w-full border h-[50px] my-7  transition ease-in-out delay-150 hover:-translate-y-1 hover:scale-110 hover:bg-gray-400 hover:text-white duration-300">
            All Logs
          </button>
          <button onClick={handleShowDrugForm} className="w-full border h-[50px] my-7  transition ease-in-out delay-150 hover:-translate-y-1 hover:scale-110 hover:bg-gray-400 hover:text-white duration-300">
            New Log
          </button>
        </div>
        <div className="space"></div>
        <div className="">
        {showDrugForm && <DrugForm />}
        {allDrugs && <Enforcer />}
        </div>
        <div className="space-2"></div>
      </div>
    </div>
  );
};

export default Confiscator;
