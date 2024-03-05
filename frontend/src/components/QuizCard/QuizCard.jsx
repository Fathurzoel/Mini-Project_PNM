// QuizCard.jsx

import React, { useState } from 'react';
import { Button, Modal, Form, Input } from 'antd';
import PropTypes from 'prop-types';

const QuizCard = ({ quiz, onDelete, onEdit }) => {
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [form] = Form.useForm();

  // Menampilkan modal edit
  const showModal = () => {
    form.setFieldsValue({
      title: quiz.title,
      description: quiz.description,
    });
    setIsModalVisible(true);
  };

  // Menutup modal
  const handleCancel = () => {
    setIsModalVisible(false);
  };

  // Menangani penghapusan kuis
  const handleDelete = () => {
    onDelete(quiz.id);
  };

  // Menangani penyuntingan kuis
  const handleEdit = () => {
    form.validateFields()
      .then(values => {
        // Memperbarui data kuis
        const updatedQuiz = {
          id: quiz.id,
          title: values.title,
          description: values.description,
        };

        // Memanggil callback onEdit
        onEdit(updatedQuiz);

        // Menutup modal
        setIsModalVisible(false);
      })
      .catch(error => {
        console.error('Validasi gagal:', error);
      });
  };

  return (
    <div>
      <div>
        <h3>{quiz.title}</h3>
        <p>{quiz.description}</p>
        <Button onClick={showModal}>Edit</Button>
        <Button onClick={handleDelete}>Delete</Button>
      </div>

      {/* Modal Edit Kuis */}
      <Modal
        title="Edit Kuis"
        visible={isModalVisible}
        onOk={handleEdit}
        onCancel={handleCancel}
      >
        <Form
          form={form}
          labelCol={{ span: 8 }}
          wrapperCol={{ span: 16 }}
          onFinish={handleEdit}
        >
          <Form.Item
            label="Judul"
            name="title"
            rules={[{ required: true, message: 'Harap masukkan judul!' }]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            label="Deskripsi"
            name="description"
            rules={[{ required: true, message: 'Harap masukkan deskripsi!' }]}
          >
            <Input.TextArea />
          </Form.Item>

          <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
            <Button type="primary" htmlType="submit">
              Simpan Perubahan
            </Button>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

// Menambahkan validasi prop untuk quiz
QuizCard.propTypes = {
  quiz: PropTypes.shape({
    id: PropTypes.string.isRequired,
    title: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
    // ...tambahkan validasi prop lainnya
  }).isRequired,
  onDelete: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
};

export default QuizCard;
