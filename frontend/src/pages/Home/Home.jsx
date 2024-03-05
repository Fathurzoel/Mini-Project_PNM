// Home.jsx

import React, { useState, useEffect } from 'react';
import { Button, Card, notification } from 'antd';
import Header from "@components/Header.jsx";
import QuizCard from "@components/QuizCard.jsx";// Sesuaikan dengan path sesuai struktur Anda
import './Home.css';
const Home = () => {
  const [activeQuizzes, setActiveQuizzes] = useState([]);
  const [completedQuizzes, setCompletedQuizzes] = useState([]);

  useEffect(() => {

    // Mock data for testing
    const mockActiveQuizzes = [
      { id: 1, title: 'Active Quiz 1', description: 'Description 1' },
      { id: 2, title: 'Active Quiz 2', description: 'Description 2' },
    ];

    const mockCompletedQuizzes = [
      { id: 3, title: 'Completed Quiz 1', description: 'Description 3' },
      { id: 4, title: 'Completed Quiz 2', description: 'Description 4' },
    ];

    setActiveQuizzes(mockActiveQuizzes);
    setCompletedQuizzes(mockCompletedQuizzes);
  }, []);

  const handleQuizDelete = (quizId) => {

    // Update the state after deletion
    setActiveQuizzes(activeQuizzes.filter(quiz => quiz.id !== quizId));
    setCompletedQuizzes(completedQuizzes.filter(quiz => quiz.id !== quizId));

    // Show a notification
    notification.success({
      message: 'Quiz Deleted',
      description: 'The quiz has been deleted successfully.',
    });
  };

  const handleQuizEdit = (updatedQuiz) => {

    // Update the state after editing
    setActiveQuizzes(activeQuizzes.map(quiz => (quiz.id === updatedQuiz.id ? updatedQuiz : quiz)));
    setCompletedQuizzes(completedQuizzes.map(quiz => (quiz.id === updatedQuiz.id ? updatedQuiz : quiz)));

    // Show a notification
    notification.success({
      message: 'Quiz Updated',
      description: 'The quiz has been updated successfully.',
    });
  };

  return (
    <Header>
    <div>
      <h1>Home</h1>

      {/* Active Quizzes */}
      <div>
        <h2>Active Quizzes</h2>
        {activeQuizzes.map(quiz => (
          <QuizCard key={quiz.id} quiz={quiz} onDelete={handleQuizDelete} onEdit={handleQuizEdit} />
        ))}
      </div>

      {/* Completed Quizzes */}
      <div>
        <h2>Completed Quizzes</h2>
        {completedQuizzes.map(quiz => (
          <Card key={quiz.id}>
            <h3>{quiz.title}</h3>
            <p>{quiz.description}</p>
          </Card>
        ))}
      </div>

      {/* Example of using Button */}
      <Button type="primary" onClick={() => console.log('Button clicked')}>
        Click me
      </Button>
    </div>
    </Header>
  );
};

export default Home;
