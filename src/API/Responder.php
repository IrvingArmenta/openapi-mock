<?php
/*
 * This file is part of Swagger Mock.
 *
 * (c) Igor Lazarev <strider2038@yandex.ru>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

namespace App\API;

use Symfony\Component\HttpFoundation\Response;

/**
 * @author Igor Lazarev <strider2038@yandex.ru>
 */
class Responder
{
    public function createResponse(int $statusCode, string $mediaType, $data): Response
    {

    }
}