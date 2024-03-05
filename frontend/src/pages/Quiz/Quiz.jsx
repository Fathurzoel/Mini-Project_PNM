import React, { useState, useEffect } from 'react';
import QuizCard from '../components/QuizCard';
import Button from '../components/Button';
import Header from '../components/Header';
import { getCompletedQuizzes, searchQuizzes } from '../services/quizService'; // Sesuaikan path sesuai struktur proyek Anda

const Quiz = () => {
  const [completedQuizzes, setCompletedQuizzes] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);

  useEffect(() => {
    fetchCompletedQuizzes();
  }, []);

  const fetchCompletedQuizzes = async () => {
    try {
      const quizzes = await getCompletedQuizzes();
      setCompletedQuizzes(quizzes);
      setSearchResults(quizzes); // Setel hasil pencarian awal ke semua kuis yang selesai
    } catch (error) {
      console.error('Error fetching completed quizzes:', error);
    }
  };

  const handleSearch = () => {
    const results = searchQuizzes(completedQuizzes, searchTerm);
    setSearchResults(results);
  };

  return (
    <div>
      <Header />
      <h2>Completed Quizzes</h2>
      <div>
        <input
          type="text"
          placeholder="Search by title"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
        <Button onClick={handleSearch}>Search</Button>
      </div>
      <div className="quiz-list">
        {searchResults.map((quiz) => (
          <QuizCard key={quiz.id} quiz={quiz} />
        ))}
      </div>
    </div>
  );
};

export default Quiz;
