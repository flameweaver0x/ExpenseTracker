import { useState, useEffect } from 'react';
import axios from 'axios';

const useFetchTransactions = (url) => {
    const [data, setData] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState(null);

    useEffect(() => {
        // Check if data is already in sessionStorage/cache
        const cachedData = sessionStorage.getItem(url);
        if (cachedData) {
            setData(JSON.parse(cachedData));
        } else {
            setIsLoading(true);
            axios.get(url)
                .then(response => {
                    // Cache data in sessionStorage
                    sessionStorage.setItem(url, JSON.stringify(response.data));
                    setData(response.data);
                })
                .catch(err => {
                    setError(err);
                })
                .finally(() => {
                    setIsLoading(false);
                });
        }
    }, [url]);

    return { data, isLoading, error };
};