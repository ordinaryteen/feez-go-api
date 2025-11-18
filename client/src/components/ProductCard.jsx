import React, { useState } from 'react';
import { Box, Image, Badge, Button, Stack, Flex, useToast } from '@chakra-ui/react';

const ProductCard = ({ product }) => {
  const [isLoading, setIsLoading] = useState(false);
  const toast = useToast();

  const handleAddToCart = async () => {
    const token = localStorage.getItem('token');
    
    if (!token) {
      toast({
        title: 'Login dulu, Bos!',
        description: 'Anda harus login untuk belanja.',
        status: 'warning',
        duration: 3000,
        isClosable: true,
      });
      return;
    }

    setIsLoading(true);

    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/cart`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}` 
        },
        body: JSON.stringify({
          product_id: product.id,
          quantity: 1 
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || 'Gagal menambahkan ke keranjang');
      }

      // 3. Sukses
      toast({
        title: 'Berhasil!',
        description: `${product.name} masuk keranjang.`,
        status: 'success',
        duration: 2000,
        isClosable: true,
      });

    } catch (error) {
      toast({
        title: 'Gagal',
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
    <Box 
      bg="gray.800" 
      maxW="sm" 
      borderWidth="1px" 
      borderRadius="lg" 
      borderColor="whiteAlpha.200"
      overflow="hidden"
      _hover={{ transform: 'scale(1.02)', transition: '0.2s', borderColor: 'brand.500' }}
    >
      <Image 
        src={`https://placehold.co/400x300/1a202c/34b27b?text=${product.name}`} 
        alt={product.name} 
        objectFit="cover"
        h="200px"
        w="100%"
      />

      <Box p="6">
        <Flex alignItems="baseline">
          <Badge borderRadius="full" px="2" colorScheme="green" bg="brand.500" color="white">
            New
          </Badge>
          <Box
            color="gray.400"
            fontWeight="semibold"
            letterSpacing="wide"
            fontSize="xs"
            textTransform="uppercase"
            ml="2"
          >
            Stock: {product.stock_tersisa}
          </Box>
        </Flex>

        <Box
          mt="1"
          fontWeight="bold"
          as="h4"
          lineHeight="tight"
          noOfLines={1}
          color="white"
          fontSize="lg"
        >
          {product.name}
        </Box>

        <Box color="gray.300">
          Rp {product.price.toLocaleString('id-ID')}
        </Box>

        <Stack mt={4} direction={'row'} spacing={4}>
          <Button 
            flex={1} 
            bg="brand.500" 
            color="white"
            _hover={{ bg: 'brand.600' }}
            size="sm"
            onClick={handleAddToCart} 
            isLoading={isLoading}
            isDisabled={product.stock_tersisa <= 0} 
          >
            {product.stock_tersisa > 0 ? 'Add to Cart' : 'Sold Out'}
          </Button>
        </Stack>
      </Box>
    </Box>
  );
};

export default ProductCard;