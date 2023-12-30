
let isValid = function () {
    if (noValues()) {
      setErrorMessage("Must enter a value");
      return false;
    }
    // if (customerIdFilter) {
    //   let regexp = /^\d+$/;
    //   if (!regexp.test(customerIdFilter)) {
    //     setErrorMessage("Customer Id can only be digits");
    //     return false;
    //   }
    //}
    // if (applicationIdFilter) {
    //   if (isNaN(parseInt(applicationIdFilter))) {
    //     setErrorMessage("ApplicationId can only be digits");
    //     return false;
    //   }
    // }
    // if (policyIdFilter) {
    //   if (isNaN(parseInt(policyIdFilter))) {
    //     setErrorMessage("PolicyId can only be digits");
    //     return false;
    //   }
    // }
    // if (carrierPolicyIdFilter) {
    //   if (carrierPolicyIdFilter.length > 100) {
    //     setErrorMessage("Cannot be longer than 100 characters");
    //     return false;
    //   }
    // }
    if ((lastNameFilter || firstNameFilter) && ssnFilter) {
      if (!Helpers.regexp_only4_ssn.test(ssnFilter)) {
        setErrorMessage(
          "SSN must be 4 digits if you search by Last Name or First Name"
        );
        return false;
      }
    }
    if (ssnFilter) {
      if (ssnFilter.length === 11 || ssnFilter.length === 9) {
        if (!Helpers.regexp_full_snn.test(ssnFilter)) {
          setErrorMessage(
            "Invalid SSN format. Must be ###-##-#### or #########"
          );
          return false;
        }
        //ok
      } else if (ssnFilter.length === 4) {
        if (!Helpers.regexp_only4_ssn.test(ssnFilter)) {
          setErrorMessage("Last 4 characters of the SSN must be 4 digits");
          return false;
        }
        if (!firstNameFilter && !lastNameFilter) {
          setErrorMessage(
            "Must enter a value for either First Name or Last Name"
          );
          return false;
        }
      } else {
        setErrorMessage("SSN must be 4 or 9 digits (dashes are optional)");
        return false;
      }
    }

    let ssnFilterLast4Characters = "";
    if (ssnFilter && ssnFilter.length === 4) {
      ssnFilterLast4Characters = ssnFilter;
    }
    let dependentInputs = [
      lastNameFilter,
      firstNameFilter,
      ssnFilterLast4Characters
    ];
    let populatedDependentInputs = dependentInputs.filter(x => x !== "");

    if (populatedDependentInputs.length === 1) {
      setErrorMessage(
        "If you use one of these inputs (Last Name, First Name, SSN), you must enter a value in at least 2 inputs : Last Name, First Name, Last 4 of SSN"
      );
      return false;
    }

    setErrorMessage("");
    return true;
  };