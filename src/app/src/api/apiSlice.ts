import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import {ApiModelsGenerateInvoiceHandlerRequest} from "../gen/client/src/models/ApiModelsGenerateInvoiceHandlerRequest"

export const api = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({ 
    baseUrl: 'http://localhost:8080',
    prepareHeaders: (headers) => {
      headers.set('Accept', 'application/pdf');
      headers.set('Content-Type', 'application/json');
      return headers;
    },
  }),
  endpoints: (builder) => ({
    createInvoice: builder.mutation<Blob, ApiModelsGenerateInvoiceHandlerRequest>({
      query: (input) => ({
        url: 'invoices',
        method: 'POST',
        body: input,
        responseHandler: (response) => response.blob(),
      }),
    }),
  }),
});

export const { useCreateInvoiceMutation } = api;
