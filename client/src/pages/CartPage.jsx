import React, { useEffect, useState } from 'react';
import { Box, Container, Heading, VStack, HStack, Text, Button, Spacer, Divider, useToast, Spinner, Center } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

const CartPage = () => {
  const [cartItems, setCartItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [isCheckingOut, setIsCheckingOut] = useState(false);
  
  const toast = useToast();
  const navigate = useNavigate();

  // 1. Fetch Keranjang pas halaman dibuka
  const fetchCart = async () => {
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/login');
      return;
    }

    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/cart`, {
        headers: { 'Authorization': `Bearer ${token}` }
      });
      
      if (response.ok) {
        const data = await response.json();
        setCartItems(data);
      } else {
        throw new Error('Gagal ambil keranjang');
      }
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchCart();
  }, []);

  // 2. Hitung Total Harga
  const totalPrice = cartItems.reduce((acc, item) => acc + (item.price_per_item * item.quantity), 0);

  // 3. Fungsi Checkout
  const handleCheckout = async () => {
    setIsCheckingOut(true);
    const token = localStorage.getItem('token');

    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/checkout`, {
        method: 'POST',
        headers: { 
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        }
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || 'Checkout gagal');
      }

      // Sukses!
      toast({
        title: 'Checkout Berhasil!',
        description: `Order ID #${data.order_id} telah dibuat. Stok berkurang.`,
        status: 'success',
        duration: 5000,
        isClosable: true,
      });

      // Kosongkan tampilan keranjang
      setCartItems([]);

    } catch (error) {
      toast({
        title: 'Gagal Checkout',
        description: error.message,
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
    } finally {
      setIsCheckingOut(false);
    }
  };

  if (loading) {
    return <Center h="50vh"><Spinner size="xl" color="brand.500" /></Center>;
  }

  return (
    <Box minH="100vh" bg="dark" color="white" py={8}>
      <Container maxW="800px">
        <Heading mb={6} color="brand.500">Your Cart</Heading>

        {cartItems.length === 0 ? (
          <Text color="gray.400">Keranjang kosong. Belanja dulu gih!</Text>
        ) : (
          <VStack spacing={4} align="stretch">
            {/* List Barang */}
            {cartItems.map((item) => (
              <Box key={item.product_id} p={4} bg="gray.800" borderRadius="md" borderWidth="1px" borderColor="whiteAlpha.200">
                <HStack>
                  <VStack align="start" spacing={0}>
                    <Text fontWeight="bold" fontSize="lg">{item.product_name}</Text>
                    <Text color="gray.400" fontSize="sm">Qty: {item.quantity}</Text>
                  </VStack>
                  <Spacer />
                  <Text fontWeight="bold" color="brand.500">
                    Rp {(item.price_per_item * item.quantity).toLocaleString('id-ID')}
                  </Text>
                </HStack>
              </Box>
            ))}

            <Divider my={4} borderColor="gray.600" />

            {/* Total & Checkout */}
            <HStack>
              <Text fontSize="xl" fontWeight="bold">Total:</Text>
              <Spacer />
              <Text fontSize="2xl" fontWeight="bold" color="white">
                Rp {totalPrice.toLocaleString('id-ID')}
              </Text>
            </HStack>

            <Button 
              size="lg" 
              bg="brand.500" 
              color="white" 
              _hover={{ bg: 'brand.600' }}
              onClick={handleCheckout}
              isLoading={isCheckingOut}
              mt={4}
            >
              Checkout Now
            </Button>
          </VStack>
        )}
      </Container>
    </Box>
  );
};

export default CartPage;