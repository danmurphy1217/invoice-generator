import React, { ReactNode } from 'react';
import styles from "./CreateInvoiceForm.module.css"

interface InvoiceFormProps {
    // the header, centered at the top of the form
    header: string;
    // children react components
    children: ReactNode;
    // the button for submitting
    button: ReactNode;
    // footer
    footer?: ReactNode;
}

const Forms: React.FC<InvoiceFormProps> = ({ header, children, button, footer }) => {
    return (
        <div className={styles.formsContainer}>
          <form className={styles.invoiceForm}>
            <h2>{header}</h2>
            <div className={styles.invoiceForm}>
              {children}
            </div>
            {button}
            {footer}
          </form>
        </div>
      );
}

export default Forms;