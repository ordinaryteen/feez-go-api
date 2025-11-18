import React, { useState } from 'react';
import { Box, Button, Container, FormControl, FormLabel, Heading, Input, VStack, useToast, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

const LoginPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  
  const toast = useToast();
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      const response = await fetch('http://localhost:8080/api/v1/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || 'Login failed');
      }

      // INI KUNCINYA: Simpan Token di browser
      localStorage.setItem('token', data.token);

      toast({
        title: 'Login Berhasil',
        status: 'success',
        duration: 3000,
        isClosable: true,
      });

      // Balik ke Home
      navigate('/');

    } catch (error) {
      toast({
        title: 'Error',
        description: error.message,
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Container maxW="md" centerContent pt={20}>
      <Box p={8} mt={8} w="100%" bg="gray.800" borderRadius="lg" borderWidth="1px" borderColor="whiteAlpha.200">
        <VStack spacing={4} align="stretch" as="form" onSubmit={handleLogin}>
          <Heading color="brand.500" textAlign="center" mb={4}>Welcome Back</Heading>

          <FormControl isRequired>
            <FormLabel color="gray.300">Email</FormLabel>
            <Input 
              type="email" 
              color="white" 
              borderColor="gray.600" 
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </FormControl>

          <FormControl isRequired>
            <FormLabel color="gray.300">Password</FormLabel>
            <Input 
              type="password" 
              color="white" 
              borderColor="gray.600"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </FormControl>

          <Button 
            type="submit" 
            bg="brand.500" 
            color="white" 
            _hover={{ bg: 'brand.600' }} 
            isLoading={isLoading}
            mt={4}
          >
            Login
          </Button>
        </VStack>
      </Box>
    </Container>
  );
};

export default LoginPage;