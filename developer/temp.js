<CustomForm onSubmit={handleSubmit}>
<ErrorModal />

<Title>Search</Title>
<Fieldset style={{ width: "500px" }}>
  <Control>
    <CustomNameLabel htmlFor="idCustomerIdFilter">
      Customer ID
    </CustomNameLabel>
    <CustomTextbox
      id="idCustomerIdFilter"
      ref={inputCustomerIdRef}
      value={customerIdFilter}
      onChange={e => {
        setCustomerIdFilter(e.target.value);
      }}
      placeholder="12345678"
    />
  </Control>
  <Control>
    <CustomNameLabel htmlFor="idMbiFilter">MBI or HICN</CustomNameLabel>
    <CustomTextbox
      id="idMbiFilter"
      value={mbiFilter}
      onChange={e => {
        setMbiFilter(e.target.value);
      }}
      placeholder="1EG4-TE5-MK73"
    />
  </Control>
  <Control orientation="inline">
    <Control style={{ flex: 2, marginRight: 10 }}>
      <CustomNameLabel htmlFor="idFirstNameFilter">
        First Name
      </CustomNameLabel>
      <CustomTextbox
        id="idFirstNameFilter"
        value={firstNameFilter}
        onChange={e => {
          setFirstNameFilter(e.target.value);
        }}
        placeholder="First name"
        style={{ width: "100%" }}
      />
    </Control>
    <Control style={{ flex: 2 }}>
      <CustomNameLabel htmlFor="idLastNameFilter">
        Last Name
      </CustomNameLabel>
      <CustomTextbox
        id="idLastNameFilter"
        value={lastNameFilter}
        onChange={e => {
          setLastNameFilter(e.target.value);
        }}
        placeholder="Last name"
        style={{ width: "100%" }}
      />
    </Control>
  </Control>
  <Control>
    <CustomNameLabel htmlFor="idApplicationIdFilter">
      Application ID
    </CustomNameLabel>
    <CustomTextbox
      id="idApplicationIdFilter"
      value={applicationIdFilter}
      onChange={e => {
        setApplicationIdFilter(e.target.value);
      }}
      placeholder="12345678"
    />
  </Control>
  <Control>
    <CustomNameLabel htmlFor="idPolicyIdFilter">
      Policy ID
    </CustomNameLabel>
    <CustomTextbox
      id="idPolicyIdFilter"
      value={policyIdFilter}
      onChange={e => {
        setPolicyIdFilter(e.target.value);
      }}
      placeholder="12345678"
    />
  </Control>
  <Control>
    <CustomNameLabel htmlFor="idCarrierPolicyIdFilter">
      Carrier Policy ID
    </CustomNameLabel>
    <CustomTextbox
      id="idCarrierPolicyIdFilter"
      value={carrierPolicyIdFilter}
      onChange={e => {
        setCarrierPolicyIdFilter(e.target.value);
      }}
      placeholder="12345678"
    />
  </Control>
  <Control>
    <CustomNameLabel htmlFor="idSsnFilter">SSN</CustomNameLabel>
    <CustomDescriptionLabel>
      Full social or last 4
    </CustomDescriptionLabel>
    <CustomTextbox
      id="idSsnFilter"
      value={ssnFilter}
      onChange={e => {
        setSsnFilter(e.target.value);
      }}
      placeholder="123-45-6789"
    />
  </Control>
  <PrimaryAndSecundaryButton>
    <Button
      id="btnClean"
      type="button"
      onClick={resetForm}
      disabled={apiCustomerSearch.isLoading}
    >
      Clear
    </Button>
    <Button
      id="btnSearch"
      type="submit"
      styleType="success"
      disabled={apiCustomerSearch.isLoading}
    >
      <Icon name="search" /> Search
    </Button>
  </PrimaryAndSecundaryButton>

  {apiCustomerSearch.isLoading && (
    <div className="animate__animated animate__flash">
      Searching...{" "}
    </div>
  )}
  {errorMessage && (
    <LightNotification
      type="danger"
      includeIcon={true}
      alwaysShowIcon={true}
      style={{ marginTop: "12px" }}
    >
      <span>{errorMessage}</span>
    </LightNotification>
  )}
</Fieldset>

</CustomForm>