// frontend/src/pages/Results/Results.js

import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import Header from '../components/Header';
import ResultCard from '../components/ResultCard';
import { getResultsByQuizID } from '../../services/resultsService';

const Results = ({ match }) => {
  const [results, setResults] = useState([]);

  useEffect(() => {
    fetchResults();
  }, []);

  const fetchResults = async () => {
    try {
      const quizID = match.params.id;
      const resultsData = await getResultsByQuizID(quizID);
      setResults(resultsData);
    } catch (error) {
      console.error('Error fetching results:', error);
    }
  };

  return (
    <div>
      <Header />
      <h2>Quiz Results</h2>
      <div className="results-list">
        {results.map((result) => (
          <ResultCard key={result.id} result={result} />
        ))}
      </div>
    </div>
  );
};

Results.propTypes = {
  match: PropTypes.object.isRequired, // Prop validation for match object
};

export default Results;
