# HHCR: hyperledger healthcare record

## Draft scenario - written by Paul:

1. Patient born, allocated NHS ID 
2. Patient registers NHS ID with GP
3. Patient requests GP consultation
4. Patient consults GP
5. GP outlines options to patient
 1. Discussion only
 2. GP performs local tests
 3. Prescription
 4. Refer patient to consultant
  1. Consultant offers appointment to patient
  2. Patient consults consultant
  3. Consultant outlines options to patient
  4. Discussion only
  5. Consultancy tests & scans
  6. Prescription
  7. Patient elects surgery
6. Patient consults pharmacist
 1. Discussion only
 2. Over the counter medicine
 3. Refer patient to GP
 4. Pharmacist dispense prescription
7. Patient dies

## System Overview

1. Patient has following functions
 1. chain_code_id = register(name,date_of_birth,gender)
 2. txId = request_consultation(date_of_appointment)
 3. query_appointment()
 4. query_information()

2. General Practitioner has following functions:
 1. txId, diagnosis_result = diagnose(patient_id)
 2. test_result = do_local_test(patient_id) 
 3. prescription = prescribe_drugs(patient_id)
 4. txId = refer_to_consultant(patient_id,date_of_appointment)
 5. query_patient()

3. Hospital has following functions:
 1. txId, diagnosis_result = diagnose(patien_id)
 2. test_result = do_clinical_test(patient_id)
 3. prescription = prescribe_drugs(patient_d) 

4. Pharmacy has following functions:
 1. sell()
 2. refer_to_gp()
 3. query()




