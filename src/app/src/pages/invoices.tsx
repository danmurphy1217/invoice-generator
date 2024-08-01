// src/pages/CreateInvoicePage.tsx
import React, { useState } from 'react';
import { useCreateInvoiceMutation } from '../api/apiSlice';
import { ApiModelsGenerateInvoiceHandlerRequest } from '../gen/client/src/models/ApiModelsGenerateInvoiceHandlerRequest';
import CreateInvoiceForm from "../components/invoices/CreateInvoiceForm";
import Spinner from "../components/reusable/Spinner";
import styles from "./Invoices.module.css";


const CreateInvoicePage: React.FC = () => {
  const [truckLink, setTruckLink] = useState("")
  const [createInvoice, { isLoading, isError }] = useCreateInvoiceMutation();
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  const isValidGarageURL = (url: string): boolean => {
    const regex = /^https?:\/\/www\.withgarage\.com\/listing\/.*-[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/;
    return regex.test(url);
  }

  const handleDownload = async (response: Blob) => {
      console.log("response", response)
      // Create a URL for the blob
      const url = window.URL.createObjectURL(response);
      console.log("url", url)
      const link = document.createElement('a');
      console.log("link", link)
      link.href = url;
      link.download = 'invoice.pdf'; // Suggested filename
      document.body.appendChild(link);
      link.click();
      
      // Clean up
      link.remove();
      window.URL.revokeObjectURL(url);
  };


  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (truckLink != "") {
      setErrorMessage(null);

      // ensure that the link is a valid garage truck link
      if (!isValidGarageURL(truckLink)) {
        setErrorMessage(`Invalid Link: ${truckLink}`);
        return
      }

      const truckId = truckLink.split('-').slice(-5).join('-')
      try {
          const msg: ApiModelsGenerateInvoiceHandlerRequest = {
            truckId: truckId,
          }
          const blob = await createInvoice(msg).unwrap();
          if (isError) {
            setErrorMessage(`Could Not Fetch Truck for Link: ${truckId}`);
            return
          } else {
            setErrorMessage(null); // Clear previous error message if successful
            console.log("handling download now");
            await handleDownload(blob)
          }
        } catch (err) {
          setErrorMessage(`Could Not Fetch Truck for Link: ${truckId}`);
          console.error(`error when creating invoice: ${err}`);
          return
        }
      }
  }

  return (
    <div className={styles.authContainer}>
      <CreateInvoiceForm
        header="Create An Invoice"
        button={
        <button type="submit" onClick={handleSubmit}>
          {isLoading ? <Spinner /> : 'Submit'}  
        </button>}
        footer={<p className={styles.smallText}>Questions? Contact <a className={styles.supportEmail} href="mailto:support@withgarage.com">support@withgarage.com</a></p>}
        >
        <div className={styles.formGroup}>
          <input type="text" id="truckId" name="Link" placeholder="Enter Link..." onChange={(e) => {setTruckLink(e.target.value)}}/>
        </div>
        {errorMessage && (
        <div className={styles.errorMessage}>
          {errorMessage}
        </div>
      )}
      </CreateInvoiceForm>
    </div>
  );
};

export default CreateInvoicePage;
